1. Microservice 1: User Management & Job Submission Service

Purpose:

Microservice 1 is the user-facing component of the job system. Its primary role is to manage user authentication, authorization, job submission, scheduling, tracking, and cancellation.

Core Responsibilities:

User Authentication & Authorization

Uses authentication methods like OAuth or JWT to verify and manage user sessions.
Manages user permissions to control access levels (e.g., admin, job manager, read-only).
Job Submission and Scheduling

Allows users to submit one-time or recurring jobs with specific schedules.
Validates job details, ensures jobs are formatted correctly, and stores job metadata in a PostgreSQL database.
Job Status Tracking

Exposes APIs that allow users to view job status in real-time.
Fetches job status from PostgreSQL, which Microservice 2 and Worker Services update based on job progress.
Job Cancellation & Rescheduling

Allows users to cancel or reschedule jobs.
If a job is canceled, Microservice 1 communicates the update to Microservice 2 to stop its processing.
Audit Logging & Monitoring

Logs all user actions, including job submissions, cancellations, and updates.
Monitors system health and user activity for auditing and debugging purposes.
Workflow:

A user submits a job request, which is validated, stored, and scheduled.
Microservice 1 responds with job ID and details, allowing the user to check on the job later.
If a job status is requested, Microservice 1 retrieves it from PostgreSQL and provides it to the user.




CREATE TABLE users (
    id SERIAL PRIMARY KEY,                           -- Unique user ID
    username VARCHAR(255) UNIQUE NOT NULL,           -- Username must be unique
    password_hash TEXT NOT NULL,                     -- Hashed password
    role VARCHAR(20) NOT NULL,                       -- User role (e.g., "admin", "user")
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),     -- Timestamp for user creation
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),     -- Timestamp for last update
    last_login TIMESTAMP                             -- Last login time, optional
);



CREATE TABLE jobs (
    id SERIAL PRIMARY KEY,                               -- Unique job ID
    user_id INT REFERENCES users(id) ON DELETE CASCADE,  -- ID of the user who submitted the job
    type VARCHAR(20) NOT NULL,                           -- Job type ("one-time" or "recurring")
    status VARCHAR(20) NOT NULL,                         -- Job status ("pending", "completed", etc.)
    scheduled_at TIMESTAMP NOT NULL,                     -- Scheduled time for the job
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),         -- Timestamp for job creation
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),         -- Last updated timestamp
    recurring_cron VARCHAR(50),                          -- Cron expression for recurring jobs
    execution_count INT DEFAULT 0,                       -- Tracks the number of executions
    last_executed_at TIMESTAMP                           -- Last execution timestamp, for recurring jobs
);


-- Index on user_id to speed up queries filtering by user
CREATE INDEX idx_jobs_user_id ON jobs (user_id);

-- Index on status to quickly find jobs by status
CREATE INDEX idx_jobs_status ON jobs (status);

-- Index on scheduled_at to efficiently find jobs by schedule
CREATE INDEX idx_jobs_scheduled_at ON jobs (scheduled_at);
