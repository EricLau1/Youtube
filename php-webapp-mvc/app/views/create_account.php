<!DOCTYPE html>
<html lang="pt-br">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>Criar conta</title>
    <link rel="stylesheet" type="text/css" href="assets/css/styles.css" />
</head>
<body>
    <div class="container">
        <header class="main-header">
            <h1> Crie sua conta </h1>
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
        <form action="/save-account" method="POST" class="form-default">
            <div>
                <label>Nickname: </label>
                <input type="text" name="nickname" placeholder="Crie um nickname"/>
            </div>
            <div>
                <label>Email: </label>
                <input type="email" name="email" placeholder="Digite seu e-mail"/>
            </div>
            <div>
                <label>Senha:</label>
                <input type="password" name="password" placeholder="Crie uma senha"/>
            </div>

                <button type="submit"> Submit </button>
        </form>
    </div>
</body>
</html>