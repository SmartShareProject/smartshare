Section "Uninstall"
  # uninstall for all users
  setShellVarContext all

  # Delete (optionally) installed files
  {{range $}}Delete $INSTDIR\{{.}}
  {{end}}
  Delete $INSTDIR\uninstall.exe

  # Delete install directory
  rmDir $INSTDIR

  # Delete start menu launcher
  Delete "$SMPROGRAMS\${APPNAME}\${APPNAME}.lnk"
  Delete "$SMPROGRAMS\${APPNAME}\Attach.lnk"
  Delete "$SMPROGRAMS\${APPNAME}\Uninstall.lnk"
  rmDir "$SMPROGRAMS\${APPNAME}"

  # Firewall - remove rules if exists
  SimpleFC::AdvRemoveRule "Smartshare incoming peers (TCP:18883)"
  SimpleFC::AdvRemoveRule "Smartshare outgoing peers (TCP:18883)"
  SimpleFC::AdvRemoveRule "Smartshare UDP discovery (UDP:18883)"

  # Remove IPC endpoint (https://github.com/smartshare/EIPs/issues/147)
  ${un.EnvVarUpdate} $0 "ETHEREUM_SOCKET" "R" "HKLM" "\\.\pipe\smartshare.ipc"

  # Remove install directory from PATH
  Push "$INSTDIR"
  Call un.RemoveFromPath

  # Cleanup registry (deletes all sub keys)
  DeleteRegKey HKLM "Software\Microsoft\Windows\CurrentVersion\Uninstall\${GROUPNAME} ${APPNAME}"
SectionEnd
