local key=KEYS[1]
local phone=ARGV[1]
local cntKey=redis.call("get",key)
local uidKey=key..":"..phone
local cnt=tonumber(redis.call("get",uidKey))
if cnt==1 then
    return -1--已经抢过了
end
local cnt=tonumber(redis.call("get",cntKey))
if cnt<=0 then
    return -2--没有了
else
    redis.call("decr",cntKey)--减掉一个库存
    redis.call("set",uidKey,1)--存一下抢到信息，
    --这里应该设置一个过期时间，但是由于不知道活动持续多久，暂时不设置了
    return 0--抢到了
end