package rest


import (
	"aoi/controllers"
	"aoi/models"

	"aoi/models/user"

    "strings"
)

type UserController struct {
	controllers.Controller
}

func (c *UserController) Read(id int64) {
    
    
	conn := c.NewConnection()

	manager := models.NewUserManager(conn)
	item := manager.Get(id)

    
    
    c.Set("item", item)
}

func (c *UserController) Index(page int, pagesize int) {
    
    
	conn := c.NewConnection()

	manager := models.NewUserManager(conn)

    var args []interface{}
    
    _email := c.Get("email")
    if _email != "" {
        args = append(args, models.Where{Column:"email", Value:_email, Compare:"like"})
    }
    _passwd := c.Get("passwd")
    if _passwd != "" {
        args = append(args, models.Where{Column:"passwd", Value:_passwd, Compare:"like"})
    }
    _name := c.Get("name")
    if _name != "" {
        args = append(args, models.Where{Column:"name", Value:_name, Compare:"="})
        
    }
    _level := c.Geti("level")
    if _level != 0 {
        args = append(args, models.Where{Column:"level", Value:_level, Compare:"="})    
    }
    _status := c.Geti("status")
    if _status != 0 {
        args = append(args, models.Where{Column:"status", Value:_status, Compare:"="})    
    }
    _elo := c.Geti("elo")
    if _elo != 0 {
        args = append(args, models.Where{Column:"elo", Value:_elo, Compare:"="})    
    }
    _image := c.Get("image")
    if _image != "" {
        args = append(args, models.Where{Column:"image", Value:_image, Compare:"like"})
    }
    _profile := c.Get("profile")
    if _profile != "" {
        args = append(args, models.Where{Column:"profile", Value:_profile, Compare:"like"})
    }
    _startdate := c.Get("startdate")
    _enddate := c.Get("enddate")
    if _startdate != "" && _enddate != "" {        
        var v [2]string
        v[0] = _startdate
        v[1] = _enddate  
        args = append(args, models.Where{Column:"date", Value:v, Compare:"between"})    
    } else if  _startdate != "" {          
        args = append(args, models.Where{Column:"date", Value:_startdate, Compare:">="})
    } else if  _enddate != "" {          
        args = append(args, models.Where{Column:"date", Value:_enddate, Compare:"<="})            
    }
    

    
    
    if page != 0 && pagesize != 0 {
        args = append(args, models.Paging(page, pagesize))
    }
    
    orderby := c.Get("orderby")
    if orderby == "" {
        if page != 0 && pagesize != 0 {
            orderby = "id desc"
            args = append(args, models.Ordering(orderby))
        }
    } else {
        orderbys := strings.Split(orderby, ",")

        str := ""
        for i, v := range orderbys {
            if i == 0 {
                str += v
            } else {
                if strings.Contains(v, "_") {                   
                    str += ", " + strings.Trim(v, " ")
                } else {
                    str += ", u_" + strings.Trim(v, " ")                
                }
            }
        }
        
        args = append(args, models.Ordering(str))
    }
    
	items := manager.Find(args)
	c.Set("items", items)

    total := manager.Count(args)
	c.Set("total", total)
}

func (c *UserController) Insert(item *models.User) {
    
    
	conn := c.NewConnection()
    
	manager := models.NewUserManager(conn)
	manager.Insert(item)

    id := manager.GetIdentity()
    c.Result["id"] = id
    item.Id = id
}

func (c *UserController) Insertbatch(item *[]models.User) {  
    if item == nil || len(*item) == 0 {
        return
    }

    rows := len(*item)
    
    
    
	conn := c.NewConnection()
    
	manager := models.NewUserManager(conn)

    for i := 0; i < rows; i++ {
	    manager.Insert(&((*item)[i]))
    }
}

func (c *UserController) Update(item *models.User) {
    
    
	conn := c.NewConnection()

	manager := models.NewUserManager(conn)
	manager.Update(item)
}

func (c *UserController) Delete(item *models.User) {
    
    
    conn := c.NewConnection()

	manager := models.NewUserManager(conn)

    
	manager.Delete(item.Id)
}

func (c *UserController) Deletebatch(item *[]models.User) {
    
    
    conn := c.NewConnection()

	manager := models.NewUserManager(conn)

    for _, v := range *item {
        
    
	    manager.Delete(v.Id)
    }
}



func (c *UserController) GetByEmail(email string) *models.User {
    
    conn := c.NewConnection()

	_manager := models.NewUserManager(conn)
    
    item := _manager.GetByEmail(email)
    
    c.Set("item", item)
    
    
    
    return item
    
}


func (c *UserController) CountByEmail(email string) int {
    
    conn := c.NewConnection()

	_manager := models.NewUserManager(conn)
    
    item := _manager.CountByEmail(email)
    
    
    
    c.Set("count", item)
    
    return item
    
}


func (c *UserController) FindByLevel(level user.Level) []models.User {
    
    conn := c.NewConnection()

	_manager := models.NewUserManager(conn)
    
    item := _manager.FindByLevel(level)
    
    
    c.Set("items", item)
    
    
    return item
    
}

// @Put()
func (c *UserController) UpdateImageById(image string ,id int64) {
    
    conn := c.NewConnection()

	_manager := models.NewUserManager(conn)
    
    _manager.UpdateImageById(image, id)
    
}


// @Put()
func (c *UserController) UpdateEmail(email string, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewUserManager(conn)
	_manager.UpdateEmail(email, id)
}

// @Put()
func (c *UserController) UpdatePasswd(passwd string, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewUserManager(conn)
	_manager.UpdatePasswd(passwd, id)
}

// @Put()
func (c *UserController) UpdateName(name string, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewUserManager(conn)
	_manager.UpdateName(name, id)
}

// @Put()
func (c *UserController) UpdateLevel(level int, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewUserManager(conn)
	_manager.UpdateLevel(level, id)
}

// @Put()
func (c *UserController) UpdateStatus(status int, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewUserManager(conn)
	_manager.UpdateStatus(status, id)
}

// @Put()
func (c *UserController) UpdateElo(elo models.Double, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewUserManager(conn)
	_manager.UpdateElo(elo, id)
}

// @Put()
func (c *UserController) UpdateImage(image string, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewUserManager(conn)
	_manager.UpdateImage(image, id)
}

// @Put()
func (c *UserController) UpdateProfile(profile string, id int64) {
    
    
	conn := c.NewConnection()

	_manager := models.NewUserManager(conn)
	_manager.UpdateProfile(profile, id)
}






