const express = require('express')
const router = express.Router()
const db = require('../config/mysql')



router.get('/',(req,res,next)=>{
   let sql = "SELECT * FROM mahasiswa";
   db.query(sql,(err,result)=>{
       if (err) throw err;
       res.status(200).json({
        message: "Get method mhasiswa",
        data: result
       })
   })
})
//reqest inseert json file
router.post('/',(req,res,next)=>{
    const nim = req.body.nim;
    const nama = req.body.nama;
    const jurusan = req.body.jurusan;
    let sql = "INSERT INTO mahasiswa (nim,nama,jurusan) VALUES ('"+nim+"','"+nama+"','"+jurusan+"')";
    db.query(sql,(err,result)=>{
    if (err) throw err;
    res.status(200).json({
        message: "Berhasil ditambahkan ",
         
      })
    })
 })

 router.get('/:nim',(req,res,next)=>{
     const nim = req.params.nim;
     let sql = "SELECT * FROM mahasiswa WHERE nim LIKE "+nim;
     db.query(sql,(err,result)=>{
     if(err) throw err;
    
     if(result.length > 0){
        res.status(200).json({
            message:"Nim Mahasiswa Ditemukan",
            data: result
           })
        }else{
        res.status(200).json({
             message: "Nim Mahasiswa Tidak Ditemukan",
             data: result
             })
            }
        
       })
       
    })

//      if(nim ==='12345'){
//         res.status(200).json({
//             message: "Nim  12345"
//         })
//      }else{
//         res.status(500).json({
//             message: "Nim  lain"
//         })
//      }
//      res.status(200).json({
//         message: "Get method mhasiswa"
//     })


 router.put('/:nama',(req,res,next)=>{
     res.status(200).json({
         message:"Update data berhasil"
     })
 })
 router.delete('/:nim'),(req,res,next)=>{
     res.status(200).json({
         message:"Delete data berhasil"
     })
 }

 module.exports = router