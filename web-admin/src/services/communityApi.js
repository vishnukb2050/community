import api from './api';

export const authService = {
    sendOTP: async (mobile) => {
        const response = await api.post('/auth/send-otp', { mobile });
        return response.data;
    },

    verifyOTP: async (mobile, otp) => {
        const response = await api.post('/auth/verify-otp', { mobile, otp });
        return response.data;
    },

    logout: async () => {
        await api.post('/auth/logout');
        localStorage.removeItem('token');
    },
};

export const communityService = {
    getCommunities: async () => {
        const response = await api.get('/communities');
        return response.data;
    },

    createCommunity: async (data) => {
        const response = await api.post('/communities', data);
        return response.data;
    },
};

export const noticeService = {
    getNotices: async (communityId) => {
        const response = await api.get(`/notices?community_id=${communityId}`);
        return response.data;
    },

    createNotice: async (data) => {
        const response = await api.post('/notices', data);
        return response.data;
    },

    deleteNotice: async (id) => {
        await api.delete(`/notices/${id}`);
    },
};

export const pollService = {
    getPoll: async (communityId) => {
        const response = await api.get(`/polls?community_id=${communityId}`);
        return response.data;
    },

    createPoll: async (data) => {
        const response = await api.post('/polls', data);
        return response.data;
    },

    getPollOptions: async (pollId) => {
        const response = await api.get(`/polls/${pollId}/options`);
        return response.data;
    },
};

export const meetingService = {
    getMinutes: async (communityId) => {
        const response = await api.get(`/minutes?community_id=${communityId}`);
        return response.data;
    },

    createMinute: async (data) => {
        const response = await api.post('/minutes', data);
        return response.data;
    },
};
