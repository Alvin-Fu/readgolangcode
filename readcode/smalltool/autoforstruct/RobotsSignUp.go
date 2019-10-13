package main
func (t *RobotsSignUp) StartTime() int64{
return t.startTime
}
func (t *RobotsSignUp) SetStartTime(startTime int64){
t.startTime = startTime
}
func (t *RobotsSignUp) SvrType() int64{
return t.svrType
}
func (t *RobotsSignUp) SetSvrType(svrType int64){
t.svrType = svrType
}
func (t *RobotsSignUp) SvrId() int64{
return t.svrId
}
func (t *RobotsSignUp) SetSvrId(svrId int64){
t.svrId = svrId
}
func (t *RobotsSignUp) OnceCount() int64{
return t.onceCount
}
func (t *RobotsSignUp) SetOnceCount(onceCount int64){
t.onceCount = onceCount
}
func (t *RobotsSignUp) RobotCount() int64{
return t.robotCount
}
func (t *RobotsSignUp) SetRobotCount(robotCount int64){
t.robotCount = robotCount
}
func (t *RobotsSignUp) Interval() int64{
return t.interval
}
func (t *RobotsSignUp) SetInterval(interval int64){
t.interval = interval
}
func (t *RobotsSignUp) MatchId() string{
return t.matchId
}
func (t *RobotsSignUp) SetMatchId(matchId string){
t.matchId = matchId
}
func (t *RobotsSignUp) GameId() string{
return t.gameId
}
func (t *RobotsSignUp) SetGameId(gameId string){
t.gameId = gameId
}
