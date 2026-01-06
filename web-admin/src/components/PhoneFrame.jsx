import React from 'react';

const PhoneFrame = ({ children }) => {
    return (
        <div className="relative mx-auto border-gray-800 dark:border-gray-850 bg-gray-800 border-[14px] rounded-[2.5rem] h-[600px] w-[300px] shadow-xl">
            <div className="h-[32px] w-[3px] bg-gray-800 absolute -left-[17px] top-[72px] rounded-l-lg"></div>
            <div className="h-[46px] w-[3px] bg-gray-800 absolute -left-[17px] top-[124px] rounded-l-lg"></div>
            <div className="h-[46px] w-[3px] bg-gray-800 absolute -left-[17px] top-[178px] rounded-l-lg"></div>
            <div className="h-[64px] w-[3px] bg-gray-800 absolute -right-[17px] top-[142px] rounded-r-lg"></div>
            <div className="rounded-[2rem] overflow-hidden w-full h-full bg-slate-950 flex flex-col relative">
                {/* Status Bar */}
                <div className="h-6 flex justify-between items-center px-6 pt-2 text-[10px] text-white/70 select-none">
                    <span>12:45</span>
                    <div className="flex gap-1">
                        <svg className="w-3 h-3" fill="currentColor" viewBox="0 0 24 24"><path d="M12 20c-4.411 0-8-3.589-8-8s3.589-8 8-8 8 3.589 8 8-3.589 8-8 8zm0-14c-3.309 0-6 2.691-6 6s2.691 6 6 6 6-2.691 6-6-2.691-6-6-6z" /></svg>
                        <svg className="w-3 h-3" fill="currentColor" viewBox="0 0 24 24"><path d="M12 4l-1.41 1.41L16.17 11H4v2h12.17l-5.58 5.59L12 20l8-8z" /></svg>
                    </div>
                </div>

                {/* Dynamic Notch */}
                <div className="absolute top-0 left-1/2 -translate-x-1/2 w-24 h-5 bg-black rounded-b-xl z-50 flex items-center justify-center gap-1">
                    <div className="w-1 h-1 bg-blue-500 rounded-full animate-pulse"></div>
                    <div className="w-8 h-1 bg-white/10 rounded-full"></div>
                </div>

                {/* Content Area */}
                <div className="flex-1 overflow-y-auto scrollbar-hide">
                    {children}
                </div>

                {/* Home Indicator */}
                <div className="h-5 flex items-center justify-center">
                    <div className="w-20 h-1 bg-white/20 rounded-full"></div>
                </div>
            </div>
        </div>
    );
};

export default PhoneFrame;
