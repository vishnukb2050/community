-- Community Management System Database Schema

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Users Table
CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    mobile VARCHAR(15) UNIQUE NOT NULL,
    name VARCHAR(255),
    email VARCHAR(255),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    allow_chat_requests BOOLEAN DEFAULT true,
    hide_phone_number BOOLEAN DEFAULT false
);

-- Communities Table
CREATE TABLE communities (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(255) NOT NULL,
    owner_id UUID REFERENCES users(id) ON DELETE SET NULL,
    subscription_status VARCHAR(50) DEFAULT 'active',
    subscription_price DECIMAL(10,2) DEFAULT 1000.00,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

-- Community Members
CREATE TABLE community_members (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    community_id UUID REFERENCES communities(id) ON DELETE CASCADE,
    user_id UUID REFERENCES users(id) ON DELETE CASCADE,
    role VARCHAR(50) DEFAULT 'member', -- 'admin', 'member'
    flat_number VARCHAR(50),
    joined_at TIMESTAMP DEFAULT NOW(),
    UNIQUE(community_id, user_id)
);

-- Expenses Table
CREATE TABLE expenses (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID REFERENCES users(id) ON DELETE CASCADE,
    amount DECIMAL(10,2) NOT NULL,
    category VARCHAR(100),
    description TEXT,
    date DATE NOT NULL,
    receipt_image_url TEXT,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

-- Expense Categories (predefined)
CREATE TABLE expense_categories (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(100) UNIQUE NOT NULL,
    keywords TEXT[], -- Array of keywords for OCR matching
    created_at TIMESTAMP DEFAULT NOW()
);

-- Insert default categories
INSERT INTO expense_categories (name, keywords) VALUES
('Food & Dining', ARRAY['cafe', 'restaurant', 'pizza', 'burger', 'starbucks', 'mcdonald', 'kfc', 'diner', 'bistro', 'pub', 'bar', 'kitchen', 'bakery']),
('Groceries', ARRAY['mart', 'supermarket', 'store', 'grocery', 'food hall', 'provisions', 'wholesale', ' walmart', 'target', 'fresh', 'dairy']),
('Travel & Fuel', ARRAY['fuel', 'petrol', 'gas', 'shell', 'exxon', 'uber', 'lyft', 'taxi', 'toll', 'parking', 'station', 'bus', 'train', 'metro', 'airline', 'flight']),
('Household', ARRAY['rent', 'maintenance', 'electricity', 'water', 'utility', 'plumbing', 'cleaning', 'repair', 'hardware', 'paint', 'furniture']),
('Health & Wellness', ARRAY['pharmacy', 'chemist', 'medical', 'hospital', 'clinic', 'doctor', 'gym', 'fitness', 'yoga', 'dental', 'lab']),
('Entertainment', ARRAY['cinema', 'movie', 'netflix', 'spotify', 'games', 'ticket', 'concert', 'theater', 'subscription', 'club']);

-- Reminders Table
CREATE TABLE reminders (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID REFERENCES users(id) ON DELETE CASCADE,
    title VARCHAR(255) NOT NULL,
    amount DECIMAL(10,2),
    due_date DATE NOT NULL,
    is_recurring BOOLEAN DEFAULT false,
    recurrence_type VARCHAR(50), -- 'monthly', 'yearly', 'custom'
    status VARCHAR(50) DEFAULT 'pending', -- 'pending', 'paid', 'snoozed', 'overdue'
    snoozed_until DATE,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

-- Notes Table
CREATE TABLE notes (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID REFERENCES users(id) ON DELETE CASCADE,
    title VARCHAR(255),
    content TEXT,
    note_type VARCHAR(50) DEFAULT 'note', -- 'note', 'checklist'
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

-- Checklist Items (for notes with type='checklist')
CREATE TABLE checklist_items (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    note_id UUID REFERENCES notes(id) ON DELETE CASCADE,
    item_text VARCHAR(255) NOT NULL,
    is_checked BOOLEAN DEFAULT false,
    position INT DEFAULT 0,
    created_at TIMESTAMP DEFAULT NOW()
);

-- Documents Table
CREATE TABLE documents (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID REFERENCES users(id) ON DELETE CASCADE,
    folder VARCHAR(100), -- 'Rent Agreements', 'ID Proof', 'Bills', etc.
    file_name VARCHAR(255),
    file_url TEXT NOT NULL,
    file_size BIGINT,
    mime_type VARCHAR(100),
    uploaded_at TIMESTAMP DEFAULT NOW()
);

-- Calendar Events (Personal)
CREATE TABLE calendar_events (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID REFERENCES users(id) ON DELETE CASCADE,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    event_date DATE NOT NULL,
    event_time TIME,
    event_type VARCHAR(50) DEFAULT 'personal', -- 'personal', 'reminder', 'community'
    related_id UUID, -- Links to reminder or community event
    created_at TIMESTAMP DEFAULT NOW()
);

-- Notices/Announcements
CREATE TABLE notices (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    community_id UUID REFERENCES communities(id) ON DELETE CASCADE,
    title VARCHAR(255) NOT NULL,
    content TEXT NOT NULL,
    priority VARCHAR(50) DEFAULT 'normal', -- 'urgent', 'normal'
    created_by UUID REFERENCES users(id),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

-- Polls
CREATE TABLE polls (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    community_id UUID REFERENCES communities(id) ON DELETE CASCADE,
    question TEXT NOT NULL,
    status VARCHAR(50) DEFAULT 'live', -- 'live', 'closed'
    created_by UUID REFERENCES users(id),
    created_at TIMESTAMP DEFAULT NOW(),
    closed_at TIMESTAMP
);

CREATE TABLE poll_options (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    poll_id UUID REFERENCES polls(id) ON DELETE CASCADE,
    option_text VARCHAR(255) NOT NULL,
    vote_count INT DEFAULT 0
);

CREATE TABLE poll_votes (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    poll_id UUID REFERENCES polls(id) ON DELETE CASCADE,
    user_id UUID REFERENCES users(id),
    option_id UUID REFERENCES poll_options(id),
    voted_at TIMESTAMP DEFAULT NOW(),
    UNIQUE(poll_id, user_id)
);

-- Meeting Minutes
CREATE TABLE meeting_minutes (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    community_id UUID REFERENCES communities(id) ON DELETE CASCADE,
    title VARCHAR(255) NOT NULL,
    meeting_date DATE NOT NULL,
    content TEXT,
    pdf_url TEXT,
    created_by UUID REFERENCES users(id),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

-- Chat Conversations
CREATE TABLE conversations (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    type VARCHAR(50), -- 'private', 'community'
    community_id UUID REFERENCES communities(id) ON DELETE CASCADE,
    last_message_at TIMESTAMP,
    created_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE conversation_participants (
    conversation_id UUID REFERENCES conversations(id) ON DELETE CASCADE,
    user_id UUID REFERENCES users(id) ON DELETE CASCADE,
    joined_at TIMESTAMP DEFAULT NOW(),
    PRIMARY KEY(conversation_id, user_id)
);

CREATE TABLE messages (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    conversation_id UUID REFERENCES conversations(id) ON DELETE CASCADE,
    sender_id UUID REFERENCES users(id),
    content TEXT NOT NULL,
    sent_at TIMESTAMP DEFAULT NOW()
);

-- Notifications
CREATE TABLE notifications (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID REFERENCES users(id) ON DELETE CASCADE,
    title VARCHAR(255) NOT NULL,
    message TEXT NOT NULL,
    notification_type VARCHAR(50), -- 'reminder', 'notice', 'chat', 'poll', 'meeting'
    related_id UUID, -- Links to the related entity
    is_read BOOLEAN DEFAULT false,
    sent_at TIMESTAMP DEFAULT NOW()
);

-- OTP Storage (temporary, for authentication)
CREATE TABLE otp_codes (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    mobile VARCHAR(15) NOT NULL,
    otp_code VARCHAR(6) NOT NULL,
    expires_at TIMESTAMP NOT NULL,
    verified BOOLEAN DEFAULT false,
    created_at TIMESTAMP DEFAULT NOW()
);

-- Indexes for better performance
CREATE INDEX idx_users_mobile ON users(mobile);
CREATE INDEX idx_expenses_user_id ON expenses(user_id);
CREATE INDEX idx_expenses_date ON expenses(date);
CREATE INDEX idx_reminders_user_id ON reminders(user_id);
CREATE INDEX idx_reminders_due_date ON reminders(due_date);
CREATE INDEX idx_notes_user_id ON notes(user_id);
CREATE INDEX idx_documents_user_id ON documents(user_id);
CREATE INDEX idx_community_members_community_id ON community_members(community_id);
CREATE INDEX idx_community_members_user_id ON community_members(user_id);
CREATE INDEX idx_notices_community_id ON notices(community_id);
CREATE INDEX idx_polls_community_id ON polls(community_id);
CREATE INDEX idx_messages_conversation_id ON messages(conversation_id);
CREATE INDEX idx_notifications_user_id ON notifications(user_id);
CREATE INDEX idx_otp_mobile ON otp_codes(mobile);

-- Triggers for updated_at timestamps
CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ language 'plpgsql';

CREATE TRIGGER update_users_updated_at BEFORE UPDATE ON users
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER update_communities_updated_at BEFORE UPDATE ON communities
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER update_expenses_updated_at BEFORE UPDATE ON expenses
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER update_reminders_updated_at BEFORE UPDATE ON reminders
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER update_notes_updated_at BEFORE UPDATE ON notes
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
