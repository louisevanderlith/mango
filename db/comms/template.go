package comms

// Use Template Engine

func populatTemplate(msg Message) string {
	template := rawTemplate()
	render := template

	return render
}

func rawTemplate() string {
	result := `<!DOCTYPE html>
	<html lang="en">
	<head>
		<title>Notification</title>
	</head>
	<body>
		<h1>Notification</h1>
		<p>{{.message}}</p>
	</body>
	</html>`

	return result
}
