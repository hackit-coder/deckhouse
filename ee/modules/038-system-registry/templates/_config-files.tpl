{{- define "template-config-files-values"  }}
files:
 - templateName: registry-manager-config.yaml
   filePath: /config/config.yaml
{{- end }}

{{- define "registry-manager-config.yaml"  }}
---
leaderElection:
  namespace: d8-system
{{- end }}