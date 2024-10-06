<?php

class Site {
    private static $conn;

    public static function init($conn) {
        self::$conn = $conn;
    }

    public static function add_site($name) {

        try {
            die(self::$conn);
            $stmt = self::$conn->prepare("INSERT INTO site ($name) VALUES (?)");
            $stmt->bindParam(1, $name);
            $stmt->execute();

            return true;
        } catch (PDOException $e) {
            return false;
        }

    }

    public static function update_site($id, $name_updated) {
        try {
            $stmt = self::$conn->prepare("UPDATE site set name = $name_updated WHERE customer.id = $id");
            $stmt->bindParam(1, $name_updated);
            $stmt->execute();

            return true;
        } catch (PDOException $e) {
            return false;
        }
    }

    public static function delete_site($id) {
        try {
            $stmt = self::$conn->prepare("DELETE site WHERE customer.id = $id");
            $stmt->bindParam(1, $id);
            $stmt->execute();

            return true;
        } catch (PDOException $e) {
            return false;
        }
    }

}