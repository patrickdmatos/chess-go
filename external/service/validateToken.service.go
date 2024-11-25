package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

// ValidateTokenWithRemoteAPI valida o token remotamente usando a API de validação
func ValidateTokenWithRemoteAPI(c *fiber.Ctx) (map[string]interface{}, error) {
	// Extrair o token do cabeçalho Authorization
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return nil, errors.New("cabeçalho Authorization ausente")
	}

	// Criar a requisição HTTP para validar o token
	req, err := http.NewRequest("GET", "http://127.0.0.1:3030/UserInfos", nil)
	if err != nil {
		return nil, fmt.Errorf("erro ao criar a requisição: %w", err)
	}

	req.Header.Set("Authorization", authHeader) // Enviar o token completo no formato "Bearer <token>"

	// Enviar a requisição HTTP
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("erro ao enviar a requisição: %w", err)
	}
	defer resp.Body.Close()

	// Verificar o status da resposta
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("token inválido ou não autorizado")
	}

	// Decodificar os claims da resposta
	var claims map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&claims); err != nil {
		return nil, fmt.Errorf("erro ao decodificar a resposta: %w", err)
	}

	return claims, nil
}
