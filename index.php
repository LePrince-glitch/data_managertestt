<?php

require_once __DIR__ . '/config.php';
require_once __DIR__ . '/src/models/Customer.php';
require_once __DIR__ . '/src/models/Number.php';
require_once __DIR__ . '/src/models/Site.php';
require_once __DIR__ . '/src/models/Recharge.php';


$uri = $_SERVER['REQUEST_URI'];

echo '<pre>';
print_r($_SERVER);
echo '</pre>';

switch ($uri) {

    case '/gestion_data/index.php':
        echo '<pre>';
        print_r($_POST);
        echo '</pre>';
        break;

    case '/gestion_data/':
        include 'home.php';
        break;

    case 'add-customer':
        break;

    case 'delete-customer':
        break;

    case 'update-customer':
        break;

    case '/':
        break;

    case '/':
        break;

    default:
        # code...
        break;
}
