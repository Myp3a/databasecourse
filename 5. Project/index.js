const express = require('express')
const bodyParser = require('body-parser')
const router = express.Router()
const cors = require('cors')

var app = express()

app.use(cors())

app.use(bodyParser.urlencoded({ extended: false }))
app.use(bodyParser.json())

app.use(require('./routes'))

app.listen(8080, () => {
    console.log('server listening on port 8080')
})

module.exports = router