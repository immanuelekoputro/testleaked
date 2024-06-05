create view vw_user_subscribes AS
select
    p.id as package_id,
    p.package_name as package_name,
    u.id as user_id,
    u.name,
    case
        when us.start_date < NOW() && NOW() < us.end_date then 1
        else 0
        end as subscribe_status,
    us.start_date as subscribe_start_date,
    us.end_date as subscribe_end_date,
    p.package_duration_days as package_duration_days
from user_subscribes us
         left join users u on us.user_id = u.id
         left join packages p on us.package_id = p.id
where u.deleted_at is null and p.deleted_at is null and us.deleted_at is null;