<?php

if(!isAuth()) {
    $_SESSION['MESSAGE'] = 'Entre com sua conta para comentar.';
    redirect('/home');
    return;
}

$auth = $_SESSION['AUTH'];
$validated = app\classes\Validator::validate($_POST);

$data = [
    'user' => $auth['user_id'],
    'comment' => $validated['comment']
];

if(!(new app\models\Feedback)->create($data)) {
    $_SESSION['MESSAGE'] = 'Não foi possível comentar. :(';
}

redirect('/home');