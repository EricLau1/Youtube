<?php

if(isset($_SESSION['AUTH'])) {
    unset($_SESSION['AUTH']);
}
redirect('/home');
