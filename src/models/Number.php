<?php

class Number {
    private static $conn;

    public static function init($conn) {
        self::$conn = $conn;
    }

    public static function add_number($num, $status, $auto_recharge) {

        try {
            die(self::$conn);
            $stmt = self::$conn->prepare("INSERT INTO numero (num, statut, recharge_auto) VALUES (?, ?, ?)");
            $stmt->bindParam(1, $num);
            $stmt->bindParam(2, $status);
            $stmt->bindParam(3, $auto_recharge);
            $stmt->execute();
            return true;
        } catch (PDOException $e) {
            return false;
        }
    }

    public static function update_number($id, $name_updated) {
        try {
            $stmt = self::$conn->prepare("UPDATE number set name = $name_updated WHERE customer.id = $id");
            $stmt->bindParam(1, $name_updated);
            $stmt->execute();

            return true;
        } catch (PDOException $e) {
            return false;
        }
    }

    public static function delete_number($id) {
        try {
            $stmt = self::$conn->prepare("DELETE number WHERE customer.id = $id");
            $stmt->bindParam(1, $id);
            $stmt->execute();

            return true;
        } catch (PDOException $e) {
            return false;
        }
    }
}
