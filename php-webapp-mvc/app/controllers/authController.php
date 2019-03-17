<?php

use app\classes\Password;
use app\models\User;

$login = app\classes\Validator::validate($_POST);
$user = (new User)->findOne(['nickname', $login['nickname']]);

if(Password::verify($login['password'], $user['password'])){
    $_SESSION["AUTH"] = [
        "is_valid" => true,
        "user_id" => $user['id']
    ];

    session_regenerate_id();
    redirect("/admin");
    return;
}

redirect("/login");

