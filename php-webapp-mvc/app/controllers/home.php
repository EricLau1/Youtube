<?php

use app\models\Feedback;

$feedbacks = (new Feedback)->getFullFeedbacks();

require $view->render('home');