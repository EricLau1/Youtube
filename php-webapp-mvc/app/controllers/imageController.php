<?php

if(!isAuth()) {
    redirect("/home");
    return;
} 

use app\models\Avatar;

if(!empty($_FILES['image']['tmp_name']) && getimagesize($_FILES['image']['tmp_name'])) {
    $image = $_FILES['image']['tmp_name'];
    $name = $_FILES['image']['name'];
    $image_content = file_get_contents($image);
    $image_encoded = base64_encode($image_content);
    
    $auth = $_SESSION['AUTH'];
    $data = [
        'user' => $auth['user_id'],
        'image' =>  $image_encoded,
        'name' => $name
    ];

    $_SESSION['MESSAGE'] = 'Avatar não foi trocado.';

    if((new Avatar)->findOne(['user', $data['user']])) {

        if((new Avatar)->update($data, ['user', $data['user']])) {
            $_SESSION['MESSAGE'] = 'Avatar trocado com sucesso!';
        }

    } else {

        if((new Avatar)->create($data)) {
            $_SESSION['MESSAGE'] = 'Avatar trocado com sucesso!';
        }

    }

    redirect('/admin');
    return;
    
}     

echo "não tem arquivo.";
