<?php

if(!isAuth()) {
    redirect('/home');
    return;
}
$auth = $_SESSION['AUTH'];
$feedbacks = (new app\models\Feedback)->findAllByUser($auth['user_id']);
$userInfo = null;
if($feedbacks) {
    $userInfo = [
        "nickname" => $feedbacks[0]["nickname"],
        "email" => $feedbacks[0]["email"],
        "avatar" => $feedbacks[0]["avatar"]
    ];

} else {
    $user = (new app\models\User)->findOne(['id', $auth['user_id']]);
    $avatar = (new app\models\Avatar)->findOne(['user', $auth['user_id']]);
    $userInfo = [
        "nickname" => $user["nickname"],
        "email" => $user["email"],
        "avatar" => $avatar["image"]
    ];
}

require $view->render('admin');