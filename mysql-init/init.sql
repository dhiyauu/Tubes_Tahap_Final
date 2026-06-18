CREATE DATABASE IF NOT EXISTS userdb;
CREATE DATABASE IF NOT EXISTS orderdb;
USE orderdb;
CREATE TABLE IF NOT EXISTS orders (
    order_id INT AUTO_INCREMENT PRIMARY KEY,
    user_id INT NOT NULL,
    resi VARCHAR(100) UNIQUE NOT NULL,
    nama_barang VARCHAR(255),
    berat INT,
    dimensi VARCHAR(100),
    jenis VARCHAR(100),
    alamat_pengirim TEXT,
    alamat_penerima TEXT,
    nama_penerima VARCHAR(255),       -- 🔥 TAMBAHKAN BARIS INI
    no_telp_penerima VARCHAR(50),     -- 🔥 TAMBAHKAN BARIS INI
    status VARCHAR(50),
    eta VARCHAR(100)
);
CREATE DATABASE IF NOT EXISTS paymentdb;
CREATE DATABASE IF NOT EXISTS trackingdb;
CREATE DATABASE IF NOT EXISTS gudangdb;
CREATE DATABASE IF NOT EXISTS courierdb;
CREATE DATABASE IF NOT EXISTS reportdb;