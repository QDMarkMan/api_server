package service

import (
	"sync"

	"github.com/demos/api_server/model"
)

// ListUser service
func ListUser(username string, offset, limit int) ([]*model.UserInfo, uint64, error) {
	// 定义中间存储容器
	infos := make([]*model.UserInfo, 0)
	users, count, err := model.ListUser(username, offset, limit)
	if err != nil {
		return nil, count, err
	}
	ids := []uint64{}

	for _, user := range users {
		ids = append(ids, user.ID)
	}

	wg := sync.WaitGroup{}
	userList := model.UserList{
		Lock:  new(sync.Mutex),
		IDMap: make(map[uint64]*model.UserInfo, len(users)),
	}
	errChan := make(chan error, 1)
	finished := make(chan bool, 1)
	for _, u := range users {
		wg.Add(1)
		go func(u *model.UserModel) {
			defer wg.Done()
			userList.Lock.Lock()
			defer userList.Lock.Unlock()
			userList.IDMap[u.ID] = &model.UserInfo{
				ID:        u.ID,
				Username:  u.Username,
				Password:  u.Password,
				CreatedAt: u.CreatedAt.Format("2020-03-08 15:04:05"),
				UpdatedAt: u.UpdatedAt.Format("2020-03-08 15:04:05"),
			}
		}(u)
	}
	go func() {
		wg.Wait()
		close(finished)
	}()
	// TODO: 这是个什么用法 ？
	select {
	case <-finished:
	case err := <-errChan:
		return nil, count, err
	}
	for _, id := range ids {
		infos = append(infos, userList.IDMap[id])
	}
	return infos, count, err
}
