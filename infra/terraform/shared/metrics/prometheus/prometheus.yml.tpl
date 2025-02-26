# my global config
global:
  scrape_interval: 15s # Set the scrape interval to every 15 seconds. Default is every 1 minute.
  evaluation_interval: 15s # Evaluate rules every 15 seconds. The default is every 1 minute.
  # scrape_timeout is set to the global default (10s).

# Alertmanager configuration
alerting:
  alertmanagers:
    - static_configs:
        - targets:
          # - alertmanager:9093

# Load rules once and periodically evaluate them according to the global 'evaluation_interval'.
rule_files:
  # - "first_rules.yml"
  # - "second_rules.yml"

# A scrape configuration containing exactly one endpoint to scrape:
# Here it's Prometheus itself.
scrape_configs:
  # The job name is added as a label `job=<job_name>` to any timeseries scraped from this config.
%{ for job in jobs ~}
  - job_name: "${job.name}"
    static_configs:
      - targets:
%{ for target in job.targets ~}
        - ${target}
%{ endfor ~}
    relabel_configs:
%{ for relabel in job.relabel ~}
      - source_labels: ${jsonencode(relabel.source_labels)}
        regex: '${relabel.regex}'
        target_label: '${relabel.target_label}'
        replacement: '${relabel.replacement}'
%{ endfor ~}
%{ endfor ~}
