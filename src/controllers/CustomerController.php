<?php 
require(__DIR__ . "/src/models/Customer.php");
require __DIR__ . "/";
class CustomerController {
    
    public static function add_customer($name) {

        Customer::add_customer($name);

    }
}
