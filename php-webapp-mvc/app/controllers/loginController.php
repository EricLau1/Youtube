<?php

if(isAuth()) {
    redirect('/admin');
    return;
}

require $view->render("login");