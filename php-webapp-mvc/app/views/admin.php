<!DOCTYPE html>
<html lang="pt-br">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>Admin</title>
    <link rel="stylesheet" type="text/css" href="assets/css/styles.css" />
</head>
<body>
    <div class="container">
        <header class="main-header">
            <h1> Admin </h1>
        </header>
        <nav class="menu">
            <ul>
                <li><a href="/home">Home</a></li>
            </ul>
            <ul>
                <li><a href="/create-account">Criar conta</a></li>
                <li><a href="/logout">Log out</a></li>
            </ul>
        </nav>
        <br />
        <p class="message">
            <?= flash() ?>
        </p>
        <div class="profile">
            <h4>Trocar avatar?</h4>
            <div class="avatar">

                <div class="image"> 
     
                    <?php if(!empty($avatar['image'])): ?>
                        <img src="data:image/jpeg;base64,<?= $avatar['image'] ?>" width="150px" height="150px"  alt="meu-avatar"/>
                    <?php else: ?>
                        Colocar uma imagem
                    <?php endif; ?>
                </div>
                <div class="avatar-form"> 
                    <form action="/image" method="POST" enctype="multipart/form-data">
                        <input type="file" name="image" />
                        <button type="submit"> Salvar </button>
                    </form> 
                </div>
            </div>
            <div class="info">
                <h3><?= $user['nickname'] ?>, <?= $user['email'] ?></h3>
            </div>
        </div>

    </div>
</body>
</html>