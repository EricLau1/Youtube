<?php

if(!isAuth()) {
    redirect('/home');
    return;
}
$auth = $_SESSION['AUTH'];
$user = (new app\models\User)->findOne(['id', $auth['user_id']]);
$avatar = (new app\models\Avatar)->findOne(['user', $user['id']]);
require $view->render('admin');