import React, { useState, useEffect } from 'react';
import { useAuth } from '../context/AuthContext';
import { meetingService } from '../services/communityApi';
import { FaPlus, FaHandshake, FaCalendarAlt, FaFileAlt } from 'react-icons/fa';

const Meetings = () => {
    const { communityId } = useAuth();
    const [minutes, setMinutes] = useState([]);
    const [loading, setLoading] = useState(true);
    const [showModal, setShowModal] = useState(false);
    const [form, setForm] = useState({ title: '', content: '', date: new Date().toISOString().split('T')[0] });

    useEffect(() => {
        if (communityId) loadMinutes();
    }, [communityId]);

    const loadMinutes = async () => {
        setLoading(true);
        try {
            const data = await meetingService.getMinutes(communityId);
            setMinutes(data.minutes || []);
        } catch (err) {
            console.error(err);
        } finally {
            setLoading(false);
        }
    };

    const handleCreate = async (e) => {
        e.preventDefault();
        try {
            await meetingService.createMinute({ ...form, community_id: communityId, meeting_date: form.date });
            setShowModal(false);
            setForm({ title: '', content: '', date: new Date().toISOString().split('T')[0] });
            loadMinutes();
        } catch (err) {
            alert('Failed to add minutes');
        }
    };

    return (
        <div className="space-y-6">
            <div className="flex justify-between items-center">
                <div>
                    <h1 className="text-3xl font-bold text-white">Meeting Minutes</h1>
                    <p className="text-slate-400">Keep record of all community gatherings and decisions</p>
                </div>
                <button
                    onClick={() => setShowModal(true)}
                    className="bg-purple-600 hover:bg-purple-700 text-white px-6 py-3 rounded-xl font-bold flex items-center space-x-2 shadow-lg hover:shadow-purple-600/20 transition-all"
                >
                    <FaPlus />
                    <span>Add Minutes</span>
                </button>
            </div>

            {loading ? (
                <div className="text-white">Loading...</div>
            ) : (
                <div className="grid gap-6">
                    {minutes.length === 0 ? (
                        <div className="bg-slate-900 border border-slate-800 rounded-2xl p-12 text-center">
                            <FaHandshake className="text-slate-700 text-5xl mx-auto mb-4" />
                            <p className="text-slate-500">No meeting records found.</p>
                        </div>
                    ) : (
                        minutes.map((m) => (
                            <div key={m.id} className="bg-slate-900 border border-slate-800 rounded-2xl overflow-hidden shadow-lg transition-all hover:border-slate-700">
                                <div className="bg-slate-800/50 p-4 border-b border-slate-800 flex justify-between items-center">
                                    <div className="flex items-center text-purple-400 font-medium">
                                        <FaCalendarAlt className="mr-2" /> {m.meeting_date}
                                    </div>
                                </div>
                                <div className="p-6">
                                    <h3 className="text-xl font-bold text-white mb-3">{m.title}</h3>
                                    <p className="text-slate-400 whitespace-pre-wrap line-clamp-4">{m.content}</p>
                                    <button className="mt-4 text-purple-500 hover:text-purple-400 text-sm font-bold uppercase tracking-wider">
                                        View Full Minutes
                                    </button>
                                </div>
                            </div>
                        ))
                    )}
                </div>
            )}

            {showModal && (
                <div className="fixed inset-0 bg-black/60 backdrop-blur-sm flex items-center justify-center p-4 z-50">
                    <div className="bg-slate-900 border border-slate-800 rounded-2xl max-w-2xl w-full p-8 shadow-2xl">
                        <h2 className="text-2xl font-bold text-white mb-6">Add Meeting Minutes</h2>
                        <form onSubmit={handleCreate} className="space-y-4">
                            <div className="grid grid-cols-2 gap-4">
                                <div className="col-span-2 md:col-span-1">
                                    <label className="block text-slate-400 text-sm mb-2">Meeting Title</label>
                                    <input
                                        type="text"
                                        value={form.title}
                                        onChange={(e) => setForm({ ...form, title: e.target.value })}
                                        className="w-full bg-slate-800 border border-slate-700 rounded-lg px-4 py-3 text-white focus:outline-none focus:ring-2 focus:ring-purple-500"
                                        placeholder="Monthly Maintenance Meeting"
                                        required
                                    />
                                </div>
                                <div className="col-span-2 md:col-span-1">
                                    <label className="block text-slate-400 text-sm mb-2">Date</label>
                                    <input
                                        type="date"
                                        value={form.date}
                                        onChange={(e) => setForm({ ...form, date: e.target.value })}
                                        className="w-full bg-slate-800 border border-slate-700 rounded-lg px-4 py-3 text-white focus:outline-none focus:ring-2 focus:ring-purple-500"
                                        required
                                    />
                                </div>
                            </div>
                            <div>
                                <label className="block text-slate-400 text-sm mb-2">Detailed Minutes</label>
                                <textarea
                                    value={form.content}
                                    onChange={(e) => setForm({ ...form, content: e.target.value })}
                                    rows="10"
                                    className="w-full bg-slate-800 border border-slate-700 rounded-lg px-4 py-3 text-white focus:outline-none focus:ring-2 focus:ring-purple-500"
                                    placeholder="Items discussed, decisions made, future actions..."
                                    required
                                ></textarea>
                            </div>
                            <div className="flex space-x-4 pt-4">
                                <button
                                    type="submit"
                                    className="flex-1 bg-purple-600 hover:bg-purple-700 text-white font-bold py-3 rounded-lg shadow-lg shadow-purple-600/20 transition-all"
                                >
                                    Save Minutes
                                </button>
                                <button
                                    type="button"
                                    onClick={() => setShowModal(false)}
                                    className="flex-1 bg-slate-800 hover:bg-slate-750 text-white font-bold py-3 rounded-lg transition-all"
                                >
                                    Cancel
                                </button>
                            </div>
                        </form>
                    </div>
                </div>
            )}
        </div>
    );
};

export default Meetings;
