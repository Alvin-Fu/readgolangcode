package main
func (t *Tmp) Agent() []Person{
return t.agent
}
func (t *Tmp) SetAgent(agent []Person){
t.agent = agent
}
func (t *Tmp) Pos() int{
return t.pos
}
func (t *Tmp) SetPos(pos int){
t.pos = pos
}
func (t *Tmp) Score() int32{
return t.score
}
func (t *Tmp) SetScore(score int32){
t.score = score
}
func (t *Tmp) LanguageType() string{
return t.languageType
}
func (t *Tmp) SetLanguageType(languageType string){
t.languageType = languageType
}
func (t *Tmp) InningScore() int32{
return t.inningScore
}
func (t *Tmp) SetInningScore(inningScore int32){
t.inningScore = inningScore
}
func (t *Tmp) Status() int{
return t.status
}
func (t *Tmp) SetStatus(status int){
t.status = status
}
func (t *Tmp) CurInning() int{
return t.curInning
}
func (t *Tmp) SetCurInning(curInning int){
t.curInning = curInning
}
func (t *Tmp) CurRound() int{
return t.curRound
}
func (t *Tmp) SetCurRound(curRound int){
t.curRound = curRound
}
func (t *Tmp) TableId() int{
return t.tableId
}
func (t *Tmp) SetTableId(tableId int){
t.tableId = tableId
}
func (t *Tmp) MatchId() string{
return t.matchId
}
func (t *Tmp) SetMatchId(matchId string){
t.matchId = matchId
}
func (t *Tmp) IsWeedOut() bool{
return t.isWeedOut
}
func (t *Tmp) SetIsWeedOut(isWeedOut bool){
t.isWeedOut = isWeedOut
}
func (t *Tmp) ReentryCount() int{
return t.reentryCount
}
func (t *Tmp) SetReentryCount(reentryCount int){
t.reentryCount = reentryCount
}
func (t *Tmp) RebuyCount() int{
return t.rebuyCount
}
func (t *Tmp) SetRebuyCount(rebuyCount int){
t.rebuyCount = rebuyCount
}
func (t *Tmp) AddonCount() int{
return t.addonCount
}
func (t *Tmp) SetAddonCount(addonCount int){
t.addonCount = addonCount
}
func (t *Tmp) Rank() int64{
return t.rank
}
func (t *Tmp) SetRank(rank int64){
t.rank = rank
}
func (t *Tmp) WeedOutTime() int64{
return t.weedOutTime
}
func (t *Tmp) SetWeedOutTime(weedOutTime int64){
t.weedOutTime = weedOutTime
}
func (t *Tmp) IsReady() bool{
return t.isReady
}
func (t *Tmp) SetIsReady(isReady bool){
t.isReady = isReady
}
func (t *Tmp) InningCount() int{
return t.inningCount
}
func (t *Tmp) SetInningCount(inningCount int){
t.inningCount = inningCount
}
func (t *Tmp) Seat() int{
return t.seat
}
func (t *Tmp) SetSeat(seat int){
t.seat = seat
}
func (t *Tmp) IsSeatPreserve() bool{
return t.isSeatPreserve
}
func (t *Tmp) SetIsSeatPreserve(isSeatPreserve bool){
t.isSeatPreserve = isSeatPreserve
}
func (t *Tmp) IsStand() bool{
return t.isStand
}
func (t *Tmp) SetIsStand(isStand bool){
t.isStand = isStand
}
func (t *Tmp) IsReentrying() bool{
return t.isReentrying
}
func (t *Tmp) SetIsReentrying(isReentrying bool){
t.isReentrying = isReentrying
}
func (t *Tmp) InningStartChip() int{
return t.inningStartChip
}
func (t *Tmp) SetInningStartChip(inningStartChip int){
t.inningStartChip = inningStartChip
}
func (t *Tmp) RebuyChip() int{
return t.rebuyChip
}
func (t *Tmp) SetRebuyChip(rebuyChip int){
t.rebuyChip = rebuyChip
}
func (t *Tmp) AddonChip() int{
return t.addonChip
}
func (t *Tmp) SetAddonChip(addonChip int){
t.addonChip = addonChip
}
func (t *Tmp) DelayStand() chan bool{
return t.delayStand
}
func (t *Tmp) SetDelayStand(delayStand chan bool){
t.delayStand = delayStand
}
