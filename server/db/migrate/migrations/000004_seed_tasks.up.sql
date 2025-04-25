INSERT INTO tasks (monthly, weekly, deadline, description, is_completed, user_id)
VALUES (False, True, CURRENT_TIMESTAMP, 'test task description', False, 1)
ON CONFLICT DO NOTHING;
