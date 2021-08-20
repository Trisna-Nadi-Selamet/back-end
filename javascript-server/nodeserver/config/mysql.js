let mysql = require('mysql2')

let con = mysql.createConnection({
    host:"localhost",
    user:"root",
    password:'12345',
    database:"kuliah"
})

con.connect(function(err){
    if (err) throw err;
    console.log("Koneksi Berhasil")
})

module.exports = con