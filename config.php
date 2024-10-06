<?php
require_once(__DIR__ . '/src/models/Customer.php');
require_once(__DIR__ . '/src/models/Recharge.php');
require_once(__DIR__ . '/src/models/Site.php');

$host = 'localhost';
$dbname = 'gestion_data';
$username = 'root';
$password = '';

try {
    $conn = new PDO("mysql:host=$host;dbname=$dbname", $username, $password);
    $conn->setAttribute(PDO::ATTR_ERRMODE, PDO::ERRMODE_EXCEPTION);

    Customer::init($conn);
    Site::init($conn);
    Recharge::init($conn);

} catch (PDOException $e) {
    die("Erreur de connexion : " . $e->getMessage());
}