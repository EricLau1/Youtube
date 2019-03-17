<!DOCTYPE html>
<html lang="pt-br">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>Login</title>
    <link rel="stylesheet" type="text/css" href="assets/css/styles.css" />
</head>
<body>
    <div class="container">
        <header class="main-header">
            <h1> Login </h1>
        </header>
        <nav class="menu">
            <ul>
                <li><a href="/home">Home</a></li>
            </ul>
            <ul>
                <li><a href="/create-account">Criar conta</a></li>
                <li><a href="/login">Logar</a></li>
            </ul>
        </nav>
        <br />
        <br />
        <form action="/auth" method="POST" class="form-default">
            <div>
                <label>Nickname: </label>
                <input type="text" name="nickname" placeholder="Digite seu nickname"/>
            </div>
            <div>
                <label>Senha:</label>
                <input type="password" name="password" placeholder="Digite sua senha"/>
            </div>

            <button type="submit"> Entrar </button>
        </form>
    </div>
</body>
</html>