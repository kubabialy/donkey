package parser

import (
	"github.com/kubabialy/donkey/lexer"
	"github.com/kubabialy/donkey/ast"
	"github.com/kubabialy/donkey/token"
)

type Parser struct {
    lxr *lexer.Lexer

    curToken  token.Token
    peekToken token.Token
}

func New(lxr *lexer.Lexer) *Parser {
    p := &Parser{lxr: lxr}

    p.NextToken()
    p.NextToken()

    return p
}

func (p *Parser) NextToken() {
    p.curToken = p.peekToken
    p.peekToken = p.lxr.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
    program := &ast.Program{}
    program.Statements = []ast.Statement{}

    for p.curToken.Type != token.EOF {
        stmt := p.parseStatement()
        if stmt != nil {
            program.Statements = append(program.Statements, stmt)
        }
        p.NextToken()
    }

    return program
}

func (p *Parser) parseStatement() ast.Statement {
    switch p.curToken.Type {
    case token.LET:
            return p.parseLetStatement()
    default:
        return nil
    }
}

func (p *Parser) parseLetStatement() *ast.LetStatement {
    stmt := &ast.LetStatement{Token: p.curToken}
    
    if !p.expectPeek(token.IDENT) {
        return nil
    }

    stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

    if !p.expectPeek(token.ASSIGN) {
        return nil
    }

    // TODO: we're skipping the expressions until w encounter a semicolon
    for !p.curTokenIs(token.SEMICOLON) {
        p.NextToken()
    }

    return stmt
}

func (p *Parser) curTokenIs(t token.TokenType) bool {
    return p.curToken.Type == t
}

func (p *Parser) peekTokenIs(t token.TokenType) bool {
    return p.peekToken.Type == t
}

func (p *Parser) expectPeek(t token.TokenType) bool {
    if p.peekTokenIs(t) {
        p.NextToken()
        return true
    } else {
        return false
    }
}
