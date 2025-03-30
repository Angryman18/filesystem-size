import CodeBlock from "@/components/CodeBlock";
import Image from "next/image";
import Link from "next/link";

const Code = ({ children }: { children: React.ReactNode }) => (
  <span className='bg-gray-200 px-1 rounded-md text-sm'>{children}</span>
);

export default function Home() {
  return (
    <div className='w-[90%] md:w-[75%] lg:w-[60%] xl:w-[40%] my-4 ml-auto mr-auto'>
      <div className='flex flex-col'>
        <h1 className='text-4xl font-bold'>Get The Directory Size (GTDS)</h1>

        <p className='mt-16 mb-4'>
          Get the size of your computer directory fast. We often have tons of folder, files or dependencies (<Code>node_modules</Code>) gathers on our computer and when actully wanted to know the size of our desktop takes a long time to calculate. Thus the <Code>gtds</Code> can help to calculate your directory size faster then ever. Try it out.
        </p>

        <Link
          href='https://github.com/Angryman18/filesystem-size'
          target='_blank'
          className='text-indigo-500 flex gap-x-2 items-center my-4'
        >
          <img src='/github-logo.png' alt='github' />
          <p>Github</p>
        </Link>
        <CodeBlock>curl -sL gtds.junior-dev.com/install.sh | sudo bash</CodeBlock>

        <div>
          <h1 className='md:text-3xl text-2xl font-bold md:my-8 my-6'>Heres How it works</h1>
        </div>
        <video src='/video.webm' autoFocus controls={false} autoPlay loop muted playsInline />

        {/* <div>
          <h1 className='md:text-3xl text-2xl font-bold md:my-8 my-6'>Heres How it works</h1>
        </div> */}

        <video
          src='/video2.webm'
          className='my-4'
          autoFocus
          controls={false}
          autoPlay
          loop
          muted
          playsInline
        />
      </div>

      <div className='mt-16 flex justify-center'>
        <a
          href='https://github.com/Angryman18'
          target='_blank'
          className=' font-bold text-slate-500 cursor-pointer hover:underline'
        >
          ❤️ Shyam Mahanta
        </a>
      </div>
    </div>
  );
}
