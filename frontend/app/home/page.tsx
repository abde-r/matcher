import NotLoggedInHomeComponent from "@/components/Home/NotLoggedIn";
import NotLoggedNavbar from "@/components/Navbar/NotLoggedIn";
import LoggedInNavbar from "@/components/Navbar/LoggedIn";

export default function Home() {
  return (
    <div className="justify-center items-center min-h-screen bg-gradient-to-r from-gray-300 to-gray-500">
      <LoggedInNavbar />
      <NotLoggedInHomeComponent />
    </div>
  );
}

