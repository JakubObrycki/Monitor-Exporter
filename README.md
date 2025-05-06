Monitor Exporter 
 -------------------------------
Monitoring system with Prometheus, Grafana and Alerts (Go + Docker)

Monitoring project written in Go:
- collects basic information about the system, including CPU, memory
- displays metrics in a format compatible with Promethus
- visualizes data in Grafana
- runs everything in docker containers
- sends e-mail alerts (SMTP) when thresholds are exceeded

---
## Technology
- Go (Golang) - own application exporting metrics
- Prometheus - data collection
- Grafana - visualization and alerting
- Docker + Docker Compose - running the entire system
- SMTP (Gmail) - sending alerts

---
## Alert
Once the 80% threshold is exceeded, an alert is triggered

Example of an alert email from grafana:

![image](https://github.com/user-attachments/assets/4f03b5db-947a-44cf-9f93-fe8fe0c75da6)

