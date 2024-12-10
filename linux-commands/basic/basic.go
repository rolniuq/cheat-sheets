package basic

const BasicName = "basic"

type Basic struct{}

func NewBasic() *Basic {
	return &Basic{}
}

func (b *Basic) GetCommands() map[string]string {
	return map[string]string{
		"apt-get":    "Package handling utility for Debian-based distributions.",
		"cat":        "Concatenate and display file content.",
		"cd":         "Change directory.",
		"chmod":      "Change file or directory permissions.",
		"chown":      "Change file or directory owner and group.",
		"cp":         "Copy files or directories.",
		"crontab":    "Schedule periodic background jobs.",
		"curl":       "Transfer data from or to a server.",
		"df":         "Report file system disk space usage.",
		"du":         "Estimate file space usage.",
		"echo":       "Display a line of text or a variable value.",
		"find":       "Search for files in a directory hierarchy.",
		"free":       "Display memory usage.",
		"grep":       "Search text using patterns.",
		"hostname":   "Show or set the system's host name.",
		"ifconfig":   "Configure a network interface.",
		"ip":         "Show/manipulate routing, devices, policy routing, and tunnels.",
		"iptables":   "Administration tool for IPv4 packet filtering and NAT.",
		"journalctl": "Query and display messages from the journal.",
		"kill":       "Terminate processes by PID.",
		"ls":         "List directory contents.",
		"man":        "Display the manual for a command.",
		"mkdir":      "Create a new directory.",
		"mount":      "Mount a file system.",
		"mv":         "Move or rename files or directories.",
		"nano":       "A simple text editor.",
		"netstat":    "Print network connections, routing tables, interface statistics, masquerade connections, and multicast memberships.",
		"ping":       "Send ICMP ECHO_REQUEST to network hosts.",
		"ps":         "Display information about running processes.",
		"pwd":        "Print working directory.",
		"rm":         "Remove files or directories.",
		"rmdir":      "Remove an empty directory.",
		"rsync":      "Remote file and directory synchronization.",
		"scp":        "Secure copy (remote file copy program).",
		"ssh":        "OpenSSH client (remote login program).",
		"sudo su":    "Allows us to switch to a different user and execute one or more commands in the shell without logging out from our current session.",
		"sudo":       "Execute a command as another user, typically the superuser.",
		"systemctl":  "Control the systemd system and service manager.",
		"tar":        "Archive files.",
		"top":        "Display and update sorted information about processes.",
		"touch":      "To create a file without any content.",
		"umount":     "Unmount a file system.",
		"uname":      "Print system information.",
		"unzip":      "Extract compressed files.",
		"uptime":     "Tell how long the system has been running.",
		"vi":         "A powerful text editor.",
		"wget":       "Retrieve files from the web.",
		"whoami":     "Display the current user.",
		"yum":        "Package manager for RPM-based distributions.",
		"zip":        "Package and compress (archive) files.",
	}
}
