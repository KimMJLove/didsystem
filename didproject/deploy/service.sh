

curl -XPUT -u elastic:admin123 http://127.0.0.1:9200/_template/did-log -H 'Content-Type: application/json' -d @template.json
{
  "index_patterns": ["log-*"],
  "settings": {
    "number_of_shards": 1
  },
  "mappings": {
    "properties": {
      "timestamp": {
        "type": "date"
      },
      "duration": {
        "type": "float"
      },
      "status": {
        "type": "integer"
      },
      "request": {
        "properties": {
          "method": {
            "type": "keyword"
          },
          "url": {
            "type": "text"
          },
          "headers": {
            "type": "object"
          },
          "body": {
            "type": "text"
          },
          "ip": {
            "type": "text"
          }
        }
      },
      "response": {
        "type": "text"
      },
      "client_ip": {
        "type": "text"
      },
      "server_ip": {
        "type": "text"
      }
    }
  }
}

curl -u elastic:admin123 'http://localhost:9200/_template/did-log'
curl -u elastic:admin123 'http://127.0.0.1:9200/log-2023.05.13/_search?size=10&pretty'