{{define "bootstrap"}}
<!DOCTYPE html>
<html lang="en">
	<head>
		<title>Reticulum Technology</title>
		<link
href="//maxcdn.bootstrapcdn.com/bootstrap/3.3.7/css/bootstrap.min.css"
rel="stylesheet">
	</head>
	<body>
		{{template "navbar"}}
		<div class="container-fluid">
			{{template "yield"}}
			
			{{template "footer"}}
		</div>

		<!-- jquery & Bootstrap JS -->
		<script src="//ajax.googleapis.com/ajax/libs/jquery/1.11.3/jquery.min.js">
		</script>
		<script src="//maxcdn.bootstrapcdn.com/bootstrap/3.3.7/js/bootstrap.min.js">
		</script>
	</body>
</html>
{{end}}