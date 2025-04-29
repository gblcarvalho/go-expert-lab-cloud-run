#!/bin/bash

BASE_URL="http://localhost:8000/weather"

CEPS=(
    "01001000"  # São Paulo
    "30140071"  # Belo Horizonte
    "20040002"  # Rio de Janeiro
    "70040900"  # Brasília
    "40010000"  # Salvador
    "80010010"  # Curitiba
    "99999999"  # CEP inexistente
    "123"       # CEP inválido (curto)
    "abcdefgh"  # CEP inválido (não numérico)
    "00000000"  # CEP inexistente
    "123456789" # CEP inválido (9 dígitos)
)

for CEP in "${CEPS[@]}"; do
    echo "======================================"
    echo "Testando CEP: $CEP"
    
    RESPONSE=$(curl -s -w "\nStatus: %{http_code}\n" -X GET "$BASE_URL/$CEP")
    
    echo "Resultado:"
    echo "$RESPONSE"
    echo ""
done
