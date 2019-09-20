package main
func (t *TmpPer) Agent() []Person{
return t.agent
}
func (t *TmpPer) SetAgent(agent []Person){
t.agent = agent
}
func (t *TmpPer) Pos() int{
return t.pos
}
func (t *TmpPer) SetPos(pos int){
t.pos = pos
}
func (t *TmpPer) Score() int32{
return t.score
}
func (t *TmpPer) SetScore(score int32){
t.score = score
}
func (t *TmpPer) LanguageType() string{
return t.languageType
}
func (t *TmpPer) SetLanguageType(languageType string){
t.languageType = languageType
}
func (t *TmpPer) InningScore() int32{
return t.inningScore
}
func (t *TmpPer) SetInningScore(inningScore int32){
t.inningScore = inningScore
}
func (t *TmpPer) Status() int{
return t.status
}
func (t *TmpPer) SetStatus(status int){
t.status = status
}
func (t *TmpPer) CurInning() int{
return t.curInning
}
func (t *TmpPer) SetCurInning(curInning int){
t.curInning = curInning
}
func (t *TmpPer) CurRound() int{
return t.curRound
}
func (t *TmpPer) SetCurRound(curRound int){
t.curRound = curRound
}
func (t *TmpPer) TableId() int{
return t.tableId
}
func (t *TmpPer) SetTableId(tableId int){
t.tableId = tableId
}
func (t *TmpPer) MatchId() string{
return t.matchId
}
func (t *TmpPer) SetMatchId(matchId string){
t.matchId = matchId
}
func (t *TmpPer) IsWeedOut() bool{
return t.isWeedOut
}
func (t *TmpPer) SetIsWeedOut(isWeedOut bool){
t.isWeedOut = isWeedOut
}
func (t *TmpPer) ReentryCount() int{
return t.reentryCount
}
func (t *TmpPer) SetReentryCount(reentryCount int){
t.reentryCount = reentryCount
}
func (t *TmpPer) RebuyCount() int{
return t.rebuyCount
}
func (t *TmpPer) SetRebuyCount(rebuyCount int){
t.rebuyCount = rebuyCount
}
func (t *TmpPer) AddonCount() int{
return t.addonCount
}
func (t *TmpPer) SetAddonCount(addonCount int){
t.addonCount = addonCount
}
func (t *TmpPer) Rank() int64{
return t.rank
}
func (t *TmpPer) SetRank(rank int64){
t.rank = rank
}
func (t *TmpPer) WeedOutTime() int64{
return t.weedOutTime
}
func (t *TmpPer) SetWeedOutTime(weedOutTime int64){
t.weedOutTime = weedOutTime
}
func (t *TmpPer) IsReady() bool{
return t.isReady
}
func (t *TmpPer) SetIsReady(isReady bool){
t.isReady = isReady
}
func (t *TmpPer) InningCount() int{
return t.inningCount
}
func (t *TmpPer) SetInningCount(inningCount int){
t.inningCount = inningCount
}
func (t *TmpPer) Seat() int{
return t.seat
}
func (t *TmpPer) SetSeat(seat int){
t.seat = seat
}
func (t *TmpPer) IsSeatPreserve() bool{
return t.isSeatPreserve
}
func (t *TmpPer) SetIsSeatPreserve(isSeatPreserve bool){
t.isSeatPreserve = isSeatPreserve
}
func (t *TmpPer) IsStand() bool{
return t.isStand
}
func (t *TmpPer) SetIsStand(isStand bool){
t.isStand = isStand
}
func (t *TmpPer) IsReentrying() bool{
return t.isReentrying
}
func (t *TmpPer) SetIsReentrying(isReentrying bool){
t.isReentrying = isReentrying
}
func (t *TmpPer) InningStartChip() int{
return t.inningStartChip
}
func (t *TmpPer) SetInningStartChip(inningStartChip int){
t.inningStartChip = inningStartChip
}
func (t *TmpPer) RebuyChip() int{
return t.rebuyChip
}
func (t *TmpPer) SetRebuyChip(rebuyChip int){
t.rebuyChip = rebuyChip
}
func (t *TmpPer) AddonChip() int{
return t.addonChip
}
func (t *TmpPer) SetAddonChip(addonChip int){
t.addonChip = addonChip
}
func (t *TmpPer) DelayStand() chan bool{
return t.delayStand
}
func (t *TmpPer) SetDelayStand(delayStand chan bool){
t.delayStand = delayStand
}
