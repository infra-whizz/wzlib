CREATE USER whizz WITH PASSWORD 'whizz';
INSERT INTO system.role_members (role, member, "isAdmin") VALUES ('admin', 'whizz', true);
