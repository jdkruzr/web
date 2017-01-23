{{define "yield"}}
<div class="row">
	<div class="col-md-4 col-md-offset-4">
		<div class="panel panel-primary">
			<div class="panel-heading">
				<h3 class="panel-title">Sign Up Now!</h3>
			</div>
			<div class="panel-body">
				{{template "signupForm"}}
			</div>
		</div>
	</div>
</div>
{{end}}

{{define "signupForm"}}
<form>
	<div class="form-group">
		<label for="email">Email address</label>
		<input type="email" class="form-control" name="email" id="email" placeholder="Email">
	</div>
	<div class="form-group">
		<label for="password">Password</label>
		<input type="password" class="form-control" name="email" id="password" placeholder="Password">
	</div>
	<button type="submit" class="btn btn-default">
	Sign Up
	</button>
</form>
{{end}}