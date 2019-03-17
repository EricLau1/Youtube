<!DOCTYPE html>
<html lang="pt-br">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>Home</title>
    <link rel="stylesheet" type="text/css" href="assets/css/styles.css" />
</head>
<body>
    <div class="container">
        <header class="main-header">
            <h1> Home </h1>
        </header>
        <nav class="menu">
            <ul>
                <li><a href="/home">Home</a></li>
            </ul>
            <ul>
                <li><a href="/create-account">Criar conta</a></li>
                <?php if(!isAuth()): ?>
                    <li><a href="/login">Logar</a></li>
                <?php else: ?>
                    <li><a href="/logout">Log Out</a></li>
                <?php endif; ?>
            </ul>
        </nav>
        <div class="feedbacks">
            <p class="message">
                <?= flash(); ?>
            </p>
            <div class="form-feedback">
                <form action="/comment" method="post" class="form-default">
                    <label> Participe da discussão!</label>
                    <textarea rows="8" cols="10" name="comment"></textarea>
                    <div>
                        <button type="submit">Comentar</button>
                    </div>
                </form>
            </div>
            <?php if(count($feedbacks) > 0): ?>
                <?php foreach($feedbacks as $f): ?>
                    <div class="feedback-users">
                        <div>
                            <img src="data:image/jpeg;base64,<?= $f['avatar'] ?>" width="50px" height="50px" alt="avatar do usuário" />
                        </div>
                        <p><span><?= $f['nickname'] ?>, escreveu:</span> <?= $f['comment']; ?></p>
                    </div>
                <?php endforeach; ?>
            <?php endif; ?>
        </div>
    </div>
</body>
</html>