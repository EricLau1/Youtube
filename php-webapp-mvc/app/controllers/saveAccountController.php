<?php

use app\models\User;

$_POST['password'] = app\classes\Password::hash($_POST['password']);
$validated = app\classes\Validator::validate($_POST);

if((new User)->create($validated)){
    redirect("/home");
    return;
}

echo "Erro ao salvar";