const pool = require('../config/config')
const router = require('express').Router()

router.post('/admin/add', (req, res, next) => {
    pool.getConnection((err, con) => {
        if (err) throw err

        let name = req.body.name
        let surname = req.body.surname
        let middle_name = req.body.middle_name
        let date_birth = req.body.date_birth

        con.query(
            'INSERT INTO `workers`(`name`,`surname`,`middle_name`,`date_birth`) VALUES (?,?,?,?)',
            [name, surname, middle_name, date_birth],
            (error, result) => {
                if (error) throw error

                res.setHeader('Content-Type', 'application/json')
                res.send(result)
            }
        )
        con.release()
    })
})

router.post('/admin/allow', (req, res, next) => {
    pool.getConnection((err, con) => {
        if (err) throw err

        let person_id = req.body.person_id
        let room_id = req.body.room_id
        // 0 - once
        // 1 - temp
        // 2 - perm
        let perm_type = req.body.perm_type
        let remaining = req.body.remaining

        con.query('SELECT * FROM `permissions` WHERE `person_id` = ? AND `room_id` = ?', (person_id, room_id), (error, result) => {
            if (error) throw error
            let set_new = false
            // existing perms
            if (result.length > 0) {
                let row = result[0]
                // exisitng is better
                if (row[3] > perm_type) {

                }
                else {
                    con.query('DELETE * FROM `permissions` WHERE `id` = ?', row[0], (error, result) => {
                        if (error) throw error
                    })
                    set_new = true
                }
            }
            // no perms
            if (set_new) {
                con.query(
                    'INSERT INTO `permissions`(`person_id`,`room_id`,`perm_type`,`remaining`) VALUES (?,?,?,?)',
                    [person_id, room_id, perm_type, remaining],
                    (error, result) => {
                        if (error) throw error

                        res.setHeader('Content-Type', 'application/json')
                        res.send(result)
                    }
                )
            }
            else {
                res.send("better exists")
            }
        })
        con.release()
    })
})

router.get('/worker/open/:room', (req, res, next) => {
    let room_id = req.params['room']
    let person_id = req.cookies.get("id")

    const report = (person_id, room_id, type) => {
        // 0 - unauthorized
        // 1 - overtime
        pool.getConnection((err, con) => {
            if (err) throw err
            con.query('INSERT INTO `incidents` (`person_id`,`room_id`,`time`,`type`) VALUES (?,?,?,?)', (person_id, room_id, Date.now(), type), (error, result) => {
                if (error) throw error
            })
            con.release()
        })
    }

    const log = (person_id, room_id) => {
        pool.getConnection((err, con) => {
            if (err) throw err
            con.query('INSERT INTO `logs` (`person_id`,`room_id`,`time`) VALUES (?,?,?)', (person_id, room_id, Date.now()), (error, result) => {
                if (error) throw error
            })
            con.query('SELECT curr_room FROM workers WHERE id = ?', person_id, (error,result) => {
                if (error) throw error
                let cur_room_id = result[0][0] // first row, only value
                // room tracking
                // no adjacent rooms exist
                if (cur_room_id == null) {
                    con.query('UPDATE workers SET curr_room = ? WHERE id = ?', (room_id, person_id), (error, result) => {
                        if (error) throw error
                    })
                }
                else {
                    con.query('UPDATE workers SET curr_room = ? WHERE id = ?', (null, person_id), (error, result) => {
                        if (error) throw error
                    })
                    con.query('SELECT max_time FROM rooms WHERE id = ?', room_id, (error, result) => {
                        if (error) throw error
                        let max_time = result[0][0]
                        con.query('SELECT * FROM logs WHERE person_id = ? AND room_id = ? ORDER BY time DESC', (person_id, room_id), (error, result) => {
                            if (error) throw error
                            let prev_time = result[0][0]
                            if (Date.now() - prev_time > max_time) {
                                report(person_id, room_id, 1)
                            }
                        })
                    })
                }
            })
            con.release()
        })
    }

    pool.getConnection((err, con) => {
        if (err) throw err

        con.query('SELECT * FROM `permissions` WHERE `person_id` = ? AND `room_id` = ?', (person_id, room_id), (error, result) => {
            if (error) throw error

            if (result.length == 0) {
                report(person_id, room_id, 0)
                res.send("fail")
            }
            else {
                // there should be only one permission row
                row = result[0]
                if (row[3] == 0) {
                    let remaining = row[4]
                    remaining = remaining - 1
                    // if remaining == 0, then permission gets removed
                    // so no situations when remaining < 0 should occur
                    if (remaining == 0) {
                        con.query('DELETE * FROM `permissions` WHERE `id` = ?', row[0], (error, result) => {
                            if (error) throw error
                        })
                    }
                    log(person_id, room_id)
                    res.send("ok")
                }
                else if (row[3] == 1) {
                    let remaining = row[4]
                    if (Date.now() < remaining) {
                        res.send("ok")
                    }
                    else {
                        con.query('DELETE * FROM `permissions` WHERE `id` = ?', row[0], (error, result) => {
                            if (error) throw error
                        })
                        report(person_id, room_id, 0)
                        res.send("fail")
                    }
                }
                else {
                    log(person_id, room_id)
                    res.send("ok")
                }
            }
        })
        con.release()
    })
})

router.post('/worker/request', (req, res, next) => {
    pool.getConnection((err, con) => {
        if (err) throw err

        let person_id = req.body.person_id
        let room_id = req.body.room_id
        let perm_type = req.body.perm_type
        let remaining = req.body.remaining

        con.query(
            'INSERT INTO `requests`(`person_id`,`room_id`,`perm_type`,`remaining`) VALUES (?,?,?,?)',
            [person_id, room_id, perm_type, remaining],
            (error, result) => {
                if (error) throw error

                res.setHeader('Content-Type', 'application/json')
                res.send(result)
            }
        )
        con.release()
    })
})

// for user list
// filtering and sorting should be done on frontend
router.get('/admin/workers', (req, res, next) => {
    pool.getConnection((err, con) => {
        if (err) throw err

        con.query('SELECT * FROM workers', (error, result) => {
            if (error) throw error

            res.send(result)
        })
        con.release()
    })
})

// for user activity
router.get('/admin/logs', (req, res, next) => {
    pool.getConnection((err, con) => {
        if (err) throw err

        con.query('SELECT * FROM logs', (error, result) => {
            if (error) throw error

            res.send(result)
        })
        con.release()
    })
})

router.get('/admin/incidents', (req, res, next) => {
    pool.getConnection((err, con) => {
        if (err) throw err

        con.query('SELECT * FROM incidents', (error, result) => {
            if (error) throw error

            res.send(result)
        })
        con.release()
    })
})

router.get('/admin/worker/logs/:id', (req, res, next) => {
    let person_id = req.params['id']
    pool.getConnection((err, con) => {
        if (err) throw err

        con.query('SELECT * FROM logs WHERE person_id = ?', person_id, (error, result) => {
            if (error) throw error

            res.send(result)
        })
        con.release()
    })
})

router.get('/admin/worker/incidents/:id', (req, res, next) => {
    let person_id = req.params['id']
    pool.getConnection((err, con) => {
        if (err) throw err

        con.query('SELECT * FROM logs WHERE person_id = ?', person_id, (error, result) => {
            if (error) throw error

            res.send(result)
        })
        con.release()
    })
})

router.delete('/admin/revoke/:person_id/:room_id', (req, res, next) => {
    let person_id = req.params['person_id']
    let room_id = req.params['room_id']

    pool.getConnection((err, con) => {
        if (err) throw err

        con.query('DELETE * FROM permissions WHERE person_id = ? AND room_id = ?', (person_id, room_id), (error, result) => {
            if (error) throw error

            res.send(result)
        })
        con.query('SELECT curr_room FROM workers WHERE id = ?', person_id, (error, result) => {
            if (error) throw error

            // in room
            if (result[0][0]) {
                con.query(
                    'INSERT INTO `permissions`(`person_id`,`room_id`,`perm_type`,`remaining`) VALUES (?,?,0,1)',
                    [person_id, room_id],
                    (error, result) => {
                        if (error) throw error
                    }
                )
            }
        })
        con.release()
    })
})

router.put('/admin/incidents/:id', (req, res, next) => {
    let incident_id = req.params['id']
    let resolution = req.body.resolution
    pool.getConnection((err, con) => {
        if (err) throw err

        con.query('UPDATE incidents SET resolution = ? WHERE id = ?', (resolution, incident_id), (error, result) => {
            if (error) throw error

            res.send(result)
        })
        con.release()
    })
})

module.exports = router