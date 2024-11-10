import NotLoggedNavbar from "@/components/Navbar/NotLoggedIn";

export default function About() {
    return (
        <div className="justify-center items-center min-h-screen bg-gradient-to-r from-gray-800 to-gray-900">
            <NotLoggedNavbar />
            About page
        </div>
    );
}