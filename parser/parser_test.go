package parser

import (
    "testing"
    "github.com/kubabialy/donkey/ast"
    "github.com/kubabialy/donkey/lexer"
)

func TestLetStatements(t *testing.T) {
    input := `
    let x = 5;
    let y = 10;
    let foobar = 838383;
    `

    lxr := lexer.New(input)
    p := New(lxr)
    
    program := p.ParseProgram()

    if program == nil {
        t.Fatalf("ParseProgram() returned nil")
    }

    if len(program.Statements) != 3 {
        t.Fatalf("Expected 3 statements, got %d", len(program.Statements))
    }

    tests := []struct {
        expectedIdentifier string
    }{
        {"x"},
        {"y"},
        {"foobar"},
    }

    for i, tt := range tests {
        stmt := program.Statements[i]
        if !testLetStatement(t, stmt, tt.expectedIdentifier) {
            return
        }
    }
}

func testLetStatement(t *testing.T, s ast.Statement, name string) bool {
    if s.TokenLiteral() != "let" {
        t.Errorf("s.TokenLiteral not 'let', got=%q", s.TokenLiteral())
        return false
    }

    letStmt, ok := s.(*ast.LetStatement)
    if !ok {
        t.Errorf("s not *ast.LetStatement. got=%T", s)
        return false
    }

    if letStmt.Name.TokenLiteral() != name {
        t.Errorf("letStmt.Name.TokenLiteral() not '%s', got=%s",
            name, letStmt.Name.TokenLiteral())
        return false
    }

    return true
}
