# Weather API by CEP

Este serviço consulta o clima baseado no CEP usando a API da ViaCEP e WeatherAPI.

## Como usar

### Requisição:

GET /cep-clima?cep=01153000

```bash
### Resposta:
- Sucesso (200):

```
```json
{
  "temp_C": 27.3,
  "temp_F": 81.1,
  "temp_K": 300.3
}
```

- CEP inválido (422):
  
```json
invalid zipcode
```

- CEP não encontrado (404):

```json
can not find zipcode
```

## Docker

```bash
docker build -t weatherapi .
docker run -p 8080:8080 -e WEATHER_API_KEY=your_key weatherapi
```

## Deploy Google Cloud Run

1. Crie projeto e container registry

2. Suba com:

```css
gcloud builds submit --tag gcr.io/[PROJECT-ID]/weatherapi
gcloud run deploy weatherapi --image gcr.io/[PROJECT-ID]/weatherapi --platform managed
```
