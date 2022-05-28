const router = require('express').Router()

router.use('/api', require('./access'))

module.exports = router