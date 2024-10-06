<?php

class Customer {
    private static $conn;

    public static function init($conn) {
        self::$conn = $conn;
    }

    public static function add_customer($name) {

        try {
            $stmt = self::$conn->prepare("INSERT INTO client (nom) VALUES (?)");
            $stmt->bindParam(1, $name);
            $stmt->execute();
            return true;
        } catch (PDOException $e) {
            return false;
        }
    }

    public static function get_customer($id) {
        try {
            $stmt = self::$conn->prepare("SELECT * FROM client WHERE client.id = ?");
            $stmt->bindParam(1, $id);
            $stmt->execute();
            $customer = $stmt->fetch();
            return $customer;
        } catch (PDOException $e) {
            return false;
        }
    }

    public static function update_customer($id, $name_updated) {
        try {
            $stmt = self::$conn->prepare("UPDATE customer set name = $name_updated WHERE customer.id = $id");
            $stmt->bindParam(1, $name_updated);
            $stmt->execute();

            return true;
        } catch (PDOException $e) {
            return false;
        }
    }

    public static function delete_customer($id) {
        try {
            $stmt = self::$conn->prepare("DELETE customer WHERE customer.id = $id");
            $stmt->bindParam(1, $id);
            $stmt->execute();

            return true;
        } catch (PDOException $e) {
            return false;
        }
    }
}
