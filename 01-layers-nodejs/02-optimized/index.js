'use strict'


const express = require('express')
const app = express()
const port = 3000

app.get('/linuxtips', (req, res) => {
  res.send('VAAAAAAAAI')
})

app.listen(port, () => {
  console.log(`Server rodando na porta: ${port}`)
})

