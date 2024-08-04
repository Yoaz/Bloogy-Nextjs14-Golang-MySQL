import Link from "next/link";

const NotFound = () => {
  return (
    <div>
      <h2>OOPSIE DAISY!</h2>
      <p>
        It seems you have taken a wrong turn. The page you are looking for is as
        elusive as a unicorn in a haystack. 🦄🌈
      </p>
      <Link href="/">Return Home</Link>
    </div>
  );
};

export default NotFound;
