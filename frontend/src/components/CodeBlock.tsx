"use client";

import { useEffect, useState } from "react";

const CodeBlock = ({ children }: { children: React.ReactNode }) => {
  const [isCopid, setIsCopid] = useState(false);
  const handleCopy = () => {
    if (isCopid) return;
    navigator.clipboard.writeText(children as string);
    setIsCopid(true);

    const timer = setTimeout(() => {
      setIsCopid(false);
      clearTimeout(timer);
    }, 1000);
  };

  return (
    <div className='w-full overflow-x-auto'>
      <div className='bg-[#343434] min-w-[600px] text-sm md:text-base md:min-w-auto text-slate-200 flex rounded-md p-4 font-mono tracking-wide'>
        {children}
        <div
          onClick={handleCopy}
          className='px-2 h-[20px] ml-auto text-sm rounded-sm bg-slate-900 cursor-pointer'
        >
          {isCopid ? "copid" : "copy"}
        </div>
      </div>
    </div>
  );
};

export default CodeBlock;
